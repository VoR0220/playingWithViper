package definitions

import (
	"fmt"
	"strings"
)

/*func replaceBlockVariable(toReplace string, do *Do) (string, error) {
	log.WithFields(log.Fields{
		"chain": do.Chain,
		"var":   toReplace,
	}).Debug("Correcting $block variable")
	blockHeight, err := GetBlockHeight(do)
	block := strconv.Itoa(blockHeight)
	log.WithField("=>", block).Debug("Current height is")
	if err != nil {
		return "", err
	}

	if toReplace == "$block" {
		log.WithField("=>", block).Debug("Replacement (=)")
		return block, nil
	}

	catchEr := regexp.MustCompile("\\$block\\+(\\d*)")
	if catchEr.MatchString(toReplace) {
		height := catchEr.FindStringSubmatch(toReplace)[1]
		h1, err := strconv.Atoi(height)
		if err != nil {
			return "", err
		}
		h2, err := strconv.Atoi(block)
		if err != nil {
			return "", err
		}
		height = strconv.Itoa(h1 + h2)
		log.WithField("=>", height).Debug("Replacement (+)")
		return height, nil
	}

	catchEr = regexp.MustCompile("\\$block\\-(\\d*)")
	if catchEr.MatchString(toReplace) {
		height := catchEr.FindStringSubmatch(toReplace)[1]
		h1, err := strconv.Atoi(height)
		if err != nil {
			return "", err
		}
		h2, err := strconv.Atoi(block)
		if err != nil {
			return "", err
		}
		height = strconv.Itoa(h1 - h2)
		log.WithField("=>", height).Debug("Replacement (-)")
		return height, nil
	}

	log.WithField("=>", toReplace).Debug("Replacement (unknown)")
	return toReplace, nil
}*/

func stringPreProcess(val string, do *Do) (string, error) {
	switch {
	/*case strings.HasPrefix(val, "$block"):
		return replaceBlockVariable(val, do)*/
	case strings.HasPrefix(val, "$"):
		key := strings.TrimPrefix(val, "$")
		if results, ok := do.Jobs.JobMap[key]; ok {
			index := strings.Index(key, ".")
			if index == -1 {
				return "", fmt.Errorf("Could not find results for job %v", index)
			} else {
				return results.JobVars[key[index:]], nil
			}
			return results.JobResult, nil
		}
		return "", fmt.Errorf("Could not find results for job %v", key)
	default: 
		return val, nil
	}
}


/*func preProcessLibs(libs string, do *definitions.Do) (string, error) {
	libraries, _ := PreProcess(libs, do)
	if libraries != "" {
		pairs := strings.Split(libraries, ",")
		for _, pair := range pairs {
			libAndAddr := strings.Split(pair, ":")
			libAndAddr[1] = strings.ToLower(libAndAddr[1])
			pair = strings.Join(libAndAddr, ":")
		}
		libraries = strings.Join(pairs, " ")
	}
	log.WithField("=>", libraries).Debug("Library String")
	return libraries, nil
}*/